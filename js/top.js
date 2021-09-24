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
                if (a.missed_pct === 0 && b.missed_pct === 0) {
                    return a.votes-b.votes
                } else if (a.missed_pct === b.missed_pct) {
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
            backgroundColor: 'rgba(0, 0, 0, 0.0)',
            textStyle: {
                color: 'rgba(255, 255, 255, 0.7)',
                fontSize: 13,
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                },
                backgroundColor: 'rgba(0,0,0,0.65)',
                textStyle: {
                    color: 'rgba(255,255,255,1.0)'
                },
                padding: [
                    5,  // up
                    10, // right
                    5,  // down
                    10, // left
                ],
                confine: true,
            },
            //legend: {
            //    data: ['Missed %', "Vote Weight"],
            //},
            grid: {
                left: '10%',
                right: '0%',
                top: '0%',
                bottom: '0%',
                containLabel: true
            },
            xAxis: {
                type: 'value',
                inverse: true,
                nameTextStyle: {fontSize: 15},
            },
            yAxis: {
                type: 'category',
                data: missing.monikers,
                nameTextStyle: {fontSize: 15},
                axisLine: { show: false },
                axisTick: { show: false },
                splitLine: { show: false },
            },
            series: [
                {
                    barGap: 0,
                    barCategoryGap: 0,
                    barWidth: 32,
                    barMaxWidth: 32,
                    name: '% Missed in '+params.depth+' blocks',
                    type: 'bar',
                    stack: 'total',
                    label: {
                        show: true,
                        position: 'insideRight',
                        padding: 8,
                        color: "#fff",
                    },
                    //emphasis: {
                    //    focus: 'series'
                    //},
                    data: missing.missed,
                    itemStyle: {
                        opacity: 0.9,
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
                    name: '% Impact vs. Consensus',
                    type: 'bar',
                    stack: 'total',
                    barGap: 0,
                    barCategoryGap: 0,
                    barWidth: 32,
                    barMaxWidth: 32,
                    label: {
                        show: true,
                        position: 'insideLeft',
                        padding: 8,
                        color: "#fff",
                    },
                    //emphasis: {
                    //    focus: 'series'
                    //},
                    data: missing.votes,
                    itemStyle: {
                        opacity: 0.9,
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
        function connectTop() {
            const socket = new WebSocket(wsProto + location.host + '/top/ws');
            socket.addEventListener('message', function (event) {
                missing = refresh(JSON.parse(event.data));
                option.yAxis.data = missing.monikers;
                option.series[0].data = missing.missed;
                option.series[1].data = missing.votes;
                chartDom.style.height = `${100 + 50 * missing.monikers.length}px`;
                myChart.setOption(option);
            });
            socket.onclose = function(e) {
                console.log('Socket is closed, retrying /top/ws ...', e.reason);
                setTimeout(function() {
                    connectTop();
                }, 4000);
            };
        }
        connectTop()

    } catch (e) {
        console.log(e.toString());
    }
}
