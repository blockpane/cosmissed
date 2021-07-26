async function topMissed() {
    try {
        const resp = await fetch("/params", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        const params = await resp.json()
        document.getElementById('timeframe').innerHTML = params.depth;
        document.getElementById('networkId').innerHTML = params.chain;

        const response = await fetch("/top", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        let data = await response.json()

        const chartDom = document.getElementById('bottom');
        chartDom.style.height = `${100+50*data.length}px`;
        const myChart = echarts.init(chartDom, 'shine');

        const refresh = function (d) {
            monikers = [];
            missed = [];
            votes = [];
            d.sort((a, b) => {
                if (a.missed_pct === b.missed_pct) {
                    return b.votes-a.votes
                }
                return a.missed_pct-b.missed_pct
            });
            d.forEach ((d) => {
                monikers.push(d.moniker);
                const m = d.missed_pct.toFixed(4)
                missed.push(m);
                const v = ((d.votes*100000000)/params.power).toFixed(4)
                votes.push(v);
            })
            return {monikers: monikers, missed: missed, votes: votes}
        }
        let missing = refresh(data)


        let option;
        option = {
            backgroundColor: '#fff',
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                }
            },
            legend: {
                data: ['Missed %', "Vote Weight"]
            },
            grid: {
                left: '6%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: {
                type: 'value',
                inverse: true
            },
            yAxis: {
                type: 'category',
                data: missing.monikers
            },
            series: [
                {
                    name: '% Missed in '+params.depth+' blocks',
                    type: 'bar',
                    stack: 'total',
                    label: {
                        show: true,
                        position: 'inside',
                    },
                    emphasis: {
                        focus: 'series'
                    },
                    data: missing.missed,
                    itemStyle: {
                        opacity: 0.8,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgb(0,0,0)'
                        }, {
                            offset: 1,
                            color: 'rgb(255,166,84)',
                        }])
                    },
                },
                {
                    name: '% Consensus Reduced',
                    type: 'bar',
                    stack: 'total',
                    label: {
                        show: true,
                        position: 'right'
                    },
                    emphasis: {
                        focus: 'series'
                    },
                    data: missing.votes,
                    itemStyle: {
                        opacity: 0.8,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgb(0,0,0)'
                        }, {
                            offset: 1,
                            color: 'rgb(89,71,190)'
                        }])
                    },
                },
            ]
        };

        option && myChart.setOption(option);

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }
        const socket = new WebSocket(wsProto+location.host+'/top/ws');
        socket.addEventListener('message', function (event) {
            missing = refresh(JSON.parse(event.data));
            option.yAxis.data = missing.monikers;
            option.series[0].data = missing.missed;
            option.series[1].data = missing.votes;
            chartDom.style.height = `${100+50*missing.monikers.length}px`;
            myChart.setOption(option);
        })

    } catch (e) {
        console.log(e.toString());
    }
}
