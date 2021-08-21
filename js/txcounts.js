async function txCounts() {
    const memResp = await fetch("/mem", {
        method: 'GET',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer'
    });
    const mempool = await memResp.json();
    let pool = []
    let blocks = []
    pool = mempool[0]
    blocks = mempool[1]

    const memDom = document.getElementById('mempool');
    const memChart = echarts.init(memDom);
    let memOption;
    memOption = {
        backgroundColor: 'rgb(7,7,7)',
        textStyle: {
            color: 'rgba(255, 255, 255, 0.9)'
        },
        //legend: {
        //    data: ['Confirmed Tx', 'Pending Tx'],
        //},

        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'line',
                lineStyle: {
                    color: 'rgba(136,220,3,0.37)',
                    width: 1,
                    type: 'solid'
                }
            }
        },

        title: {
            text: "Tx Count: mempool vs. blocks",
            textStyle: {
                color: "#FFF",
                fontSize: 13,
                fontWeight: "lighter",
            },
        },

        singleAxis: {
            top: 50,
            bottom: 50,
            axisTick: {},
            axisLabel: {},
            type: 'time',
            axisPointer: {
                animation: true,
                label: {
                    show: false
                }
            },
            splitLine: {
                show: true,
                lineStyle: {
                    type: 'dashed',
                    opacity: 0.2
                }
            },
            textStyle: {
                color: 'rgba(255, 255, 255, 0.9)'
            }
        },

        series: [
            {
                type: 'themeRiver',
                useUtc: true,
                itemStyle: {
                    color: {
                        type: 'linear',
                        x: 0,
                        y: 0,
                        x2: 0,
                        y2: 0.8,
                        colorStops: [{
                            offset: 0, color: '#F000D2' // color at 0% position
                        }, {
                            offset: 1, color: '#0C0057' // color at 100% position
                        }],
                        global: false // false by default
                    }
                },
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowColor: 'rgba(255, 255, 255, 0.5)',
                    },
                },
                label: {
                    show: false,
                },
                z: 2,
                animation: true,
                animationDuration: 1500,
                stateAnimation: {
                    duration: 1000,
                },
                data: pool,
            },
            {
                type: 'themeRiver',
                useUtc: true,
                animation: false,
                animationDuration: 300,
                z: 0,
                itemStyle: {
                    color: {
                        type: 'linear',
                        x: 0,
                        y: 0.5,
                        x2: 0,
                        y2: 1,
                        colorStops: [{
                            offset: 0, color: '#000' // color at 0% position
                        }, {
                            offset: 1, color: 'rgba(255, 255, 255, 0.8)' // color at 100% position
                        }],
                        global: false // false by default
                    },
                    //borderColor: 'rgba(234,148,63,0.7)',
                    borderColor: 'rgb(238,223,208)',
                    borderWidth: 0.2,
                    shadowBlur: 15,
                    shadowColor: 'rgba(255, 255, 255, 0.9)',
                },
                label: {
                    show: false,
                },
                data: blocks
            },
        ]
    };
    memOption && memChart.setOption(memOption)

    let wsProto = "ws://"
    if (location.protocol === "https:") {
        wsProto = "wss://"
    }

    function connectMem() {
        const socket = new WebSocket(wsProto + location.host + '/mem/ws');
        socket.addEventListener('message', function (event) {
            const mp = JSON.parse(event.data);
            if (mp[2] === "Pending Tx") {
                pool.shift();
                pool.push(mp);
                blocks.shift();
                blocks.push([mp[0],0,"Confirmed Tx"]);
                document.getElementById('pendingTx').innerHTML = mp[1]
            } else {
                blocks[blocks.length -1] = mp;
                document.getElementById('lastTx').innerHTML = mp[1]
            }
            const d = new Date();
            if ((d.getTime() / 1000) % 6 < 1) {
                memOption.series[0].data = pool
                memOption.series[1].data = blocks
                memChart.setOption(memOption);
            }
        });
        socket.onclose = function (e) {
            console.log('Socket is closed, retrying /mem/ws ...', e.reason);
            setTimeout(function () {
                connectMem();
            }, 4000);
        };
    }

    connectMem()
}