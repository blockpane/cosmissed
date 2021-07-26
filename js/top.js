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

        const response = await fetch("/top", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        const data = await response.json()

        const chartDom = document.getElementById('bottom');
        chartDom.style.height = `${100+50*data.length}px`;
        console.log(`${30*data.length}px`);
        const myChart = echarts.init(chartDom, 'shine');

        let monikers = [];
        let missed = [];
        let votes = [];

        data.sort((a, b) => {return a.missed_pct-b.missed_pct});
        data.forEach ((d) => {
            monikers.push(d.moniker);
            const m = d.missed_pct.toFixed(4)
            missed.push(m);
            const v = ((d.votes*100000000)/params.power).toFixed(4)
            votes.push(v);
        })

        let option;
        option = {
            //backgroundColor: '#0e0e0e',
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
                data: monikers
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
                    data: missed,
                    itemStyle: {
                        opacity: 0.8,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            //color: 'rgb(30,21,79)'
                            color: 'rgb(0,0,0)'
                        }, {
                            offset: 1,
                            color: 'rgb(255,166,84)',
                            //color: 'rgba(1, 191, 236)'
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
                    data: votes,
                    itemStyle: {
                        opacity: 0.8,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            //color: 'rgb(30,21,79)'
                            color: 'rgb(0,0,0)'
                        }, {
                            offset: 1,
                            color: 'rgb(89,71,190)'
                            //color: 'rgba(1, 191, 236)'
                        }])
                    },
                },
            ]
        };

        option && myChart.setOption(option);

    } catch (e) {
        console.log(e.toString());
    }
}
