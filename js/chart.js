async function query() {
    try {
        const response = await fetch("/chart", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        const data = await response.json()

        const chartDom = document.getElementById('main');
        const myChart = echarts.init(chartDom, 'shine');
        let option;

        option = {
            //backgroundColor: '#0e0e0e',
            backgroundColor: '#fff',
            title: {
                text: 'Missing Signatures',
                subtext: 'block time',
                left: 'center'
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    animation: false
                }
            },
            legend: {
                data: ['legend 1', 'legend 2'],
                left: 10
            },
            toolbox: {
                feature: {
                    dataZoom: {
                        yAxisIndex: 'none'
                    },
                    restore: {},
                    saveAsImage: {}
                }
            },
            axisPointer: {
                link: {xAxisIndex: 'all'}
            },
            dataZoom: [
                {
                    show: true,
                    realtime: true,
                    start: 90,
                    end: 100,
                    xAxisIndex: [0, 1],
                    backgroundColor: 'rgba(104,104,104,0.22)',
                    dataBackground: {
                        lineStyle: {
                            color: '#000',
                            width: 1,
                        }
                    }
                },
                {
                    type: 'inside',
                    backgroundColor: 'rgba(255,255,255,0.55)',
                    realtime: true,
                    start: 90,
                    end: 100,
                    xAxisIndex: [0, 1]
                }
            ],
            grid: [{
                left: 50,
                right: 50,
                height: '45%'
            }, {
                left: 50,
                right: 50,
                top: '55%',
                height: '25%'
            }],
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    axisLine: {onZero: true},
                    data: data.time,
                },
                {
                    gridIndex: 1,
                    type: 'category',
                    boundaryGap: false,
                    axisLine: {onZero: true},
                    data: data.blocks,
                    position: 'top'
                }
            ],
            yAxis: [
                {
                    name: 'Validators',
                    type: 'value',
                    //max: 500
                },
                {
                    gridIndex: 1,
                    name: 'Seconds',
                    type: 'value',
                    inverse: true,
                    min: 5,
                    textStyle: {
                        color: '#fff'
                    }

                }
            ],
            series: [
                {
                    name: 'missing validator signatures',
                    type: 'line',
                    smooth: true,
                    symbolSize: 0,
                    hoverAnimation: false,
                    data: data.missed,
                    lineStyle: {
                        color: 'rgb(180,172,222)',
                    },
                    areaStyle: {
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
                {
                    name: 'missing consensus %',
                    type: 'line',
                    smooth: false,
                    symbolSize: 0,
                    hoverAnimation: true,
                    data: data.missing_percent,
                    lineStyle: {
                        color: 'rgb(238,131,25)',
                        width: 1.0,
                        type: 'dashed',
                    },
                },
                {
                    name: 'seconds since last block',
                    type: 'line',
                    smooth: true,
                    xAxisIndex: 1,
                    yAxisIndex: 1,
                    symbolSize: 0,
                    hoverAnimation: false,
                    data: data.took,
                    areaStyle: {
                        opacity: 0.8,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgb(89,71,190)'
                        }, {
                            offset: 1,
                            //color: 'rgb(30,21,79)'
                            color: 'rgb(0,0,0)'
                        }])
                    },
                    lineStyle: {
                        color: 'rgb(180,172,222)',
                        width: 0.9,
                    },
                    label: {
                        color: '#fff',
                    },
                }
            ]
        };

        option && myChart.setOption(option);
    } catch (e) {
        console.log(e.toString());
    }
}
