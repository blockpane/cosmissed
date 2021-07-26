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
        let data = await response.json()

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

        const trackPos = function (e) {
            option.dataZoom[0].start = e.start;
            option.dataZoom[1].start = e.start;
            option.dataZoom[0].end = e.end;
            option.dataZoom[1].end = e.end;
        };
        myChart.on('dataZoom', function(e) {
            trackPos(e)
        });
        myChart.on('restore', function(e) {
            trackPos(e)
        });

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }
        const socket = new WebSocket(wsProto+location.host+'/chart/ws');
        socket.addEventListener('message', function (event) {
            const upd = JSON.parse(event.data);
            data.blocks.shift();
            data.blocks.push(upd.block);
            console.log(upd.block);
            data.time.shift();
            data.time.push(upd.time);
            data.missed.shift();
            data.missed.push(upd.missed);
            data.missing_percent.shift();
            data.missing_percent.push(upd.missing_percent);
            data.took.shift();
            data.took.push(upd.took);
            myChart.setOption(option);
        });

    } catch (e) {
        console.log(e.toString());
    }
}
