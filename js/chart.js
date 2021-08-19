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
        const missingChart = echarts.init(chartDom, 'shine');
        let option;

        option = {
            backgroundColor: 'rgba(255, 255, 255, 0.0)',
            textStyle: {
                color: 'rgba(255, 255, 255, 0.7)'
            },
            title: {
                text: 'Missing Signatures',
                left: 'center'
            },
            tooltip: {
                trigger: 'axis',
                triggerOn: "mousemove",
                axisPointer: {
                    animation: false
                },
                order: 'valueDesc',
            },
            axisPointer: {
                link: {xAxisIndex: 'all'},
                type: "cross",
            },
            dataZoom: [
                {
                    show: true,
                    realtime: true,
                    start: 97,
                    end: 100,
                    zoomOnMouseWheel: false,
                    moveOnMouseWheel: false,
                    moveOnMouseMove: false,
                    xAxisIndex: [0, 1],
                    backgroundColor: 'rgba(104,104,104,0.22)',
                    dataBackground: {
                        lineStyle: {
                            color: '#333',
                            width: 2,
                        }
                    }
                },
            ],
            grid: [{
                left: 50,
                right: 50,
                height: '57%'
            }, {
                left: 50,
                right: 50,
                top: '75%',
                height: '20%'
            }],
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    axisLine: {onZero: true},
                    data: data.time,
                    show: false,
                },
                {
                    gridIndex: 1,
                    type: 'category',
                    boundaryGap: false,
                    axisLine: {onZero: true},
                    data: data.blocks,
                    position: 'top',
                    axisLabel: {
                        color: '#fff',
                        verticalAlign: 'bottom',
                    },
                }
            ],
            yAxis: [
                {
                    name: 'Missing',
                    type: 'value',
                },
                {
                    gridIndex: 1,
                    name: 'Block Time',
                    type: 'value',
                    inverse: true,
                }
            ],
            series: [
                {
                    name: 'missing validator signatures',
                    type: 'line',
                    z: 0,
                    smooth: true,
                    symbolSize: 0,

                    // little hack that makes the chart clickable, only the "symbol" will emit a mouse click event,
                    // this will create a giant clear svg and overlay so the whole series is clickable:
                    symbol: 'image://data:image/svg;base64,PD94bWwgdmVyc2lvbj0iMS4wIj8+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjYwMCI+CiAgICA8cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNGRkYiIHN0cm9rZS13aWR0aD0iMCIgZD0ibTAsMGg0ODB2MjcwSDB6Ii8+Cjwvc3ZnPg==',
                    symbolSize: [200, 600],
                    symbolOffset: [0, "50%"],
                    showSymbol: false,

                    hoverAnimation: false,
                    data: data.missed,
                    lineStyle: {
                        color: 'rgb(180,172,222)',
                    },
                    areaStyle: {
                        opacity: 0.9,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgb(0,0,0)'
                        }, {
                            offset: 0.5,
                            color: 'rgb(94,83,173)'
                        }, {
                            offset: 1,
                            color: 'rgb(31,27,61)'
                        }]),
                        shadowColor: 'rgb(109,91,210,0.3)',
                        shadowBlur: 4,
                        shadowOffsetY: -2,
                        shadowOffsetX: 0,
                    },
                },
                {
                    name: 'missing consensus %',
                    type: 'line',
                    z: 1,
                    smooth: true,

                    symbol: 'image://data:image/svg;base64,PD94bWwgdmVyc2lvbj0iMS4wIj8+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjYwMCI+CiAgICA8cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNGRkYiIHN0cm9rZS13aWR0aD0iMCIgZD0ibTAsMGg0ODB2MjcwSDB6Ii8+Cjwvc3ZnPg==',
                    symbolSize: [200, 600],
                    symbolOffset: [0, "50%"],
                    showSymbol: false,

                    hoverAnimation: true,
                    data: data.missing_percent,
                    lineStyle: {
                        color: 'rgb(238,131,25)',
                        width: 2.0,
                        type: 'dotted',
                    },
                },
                {
                    name: 'seconds since last block',
                    type: 'line',
                    smooth: true,
                    xAxisIndex: 1,
                    yAxisIndex: 1,
                    symbol: 'image://data:image/svg;base64,PD94bWwgdmVyc2lvbj0iMS4wIj8+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjYwMCI+CiAgICA8cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNGRkYiIHN0cm9rZS13aWR0aD0iMCIgZD0ibTAsMGg0ODB2MjcwSDB6Ii8+Cjwvc3ZnPg==',
                    symbolSize: [200, 600],
                    showSymbol: false,
                    hoverAnimation: false,
                    data: data.took,
                    areaStyle: {
                        opacity: 0.9,
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                            offset: 0,
                            color: 'rgb(109,91,210)'
                        }, {
                            offset: 1,
                            color: 'rgb(0,0,0)'
                        }]),
                        shadowColor: 'rgb(0,0,0,0.3)',
                        shadowBlur: 5,
                        shadowOffsetY: 6,
                        shadowOffsetX: 0
                    },
                    lineStyle: {
                        color: 'rgb(180,172,222)',
                        width: 0.9,
                    },
                }
            ]
        };

        option && missingChart.setOption(option);

        missingChart.on('dataZoom', function (e) {
            option.dataZoom[0].start = e.start;
            option.dataZoom[0].end = e.end;
        });

        const setMissing = function (title, data, jailed){
            let monikers = [];
            let titleId = "missingWhen";
            let listId = "missing"
            if (jailed) {
                titleId = "currentJailed"
                listId = "jailed"
            }
            let missingData
            if (jailed) {
                missingData = data.jailed_unbonding;
            } else {
                missingData = data.missing
            }
            if (typeof(missingData) != "undefined") {
                for (const [key, value] of Object.entries(missingData)) {
                    monikers.push(key)
                }
            }
            if (monikers.length === 0 && (title === "Currently Missing:" || title === "Jailed (unbonding):")) {
                title = ""
            }
            const missingWhen = document.getElementById(titleId);
            missingWhen.innerHTML = title+'<br/>&nbsp;<br/>&nbsp;'
            const missing = document.getElementById(listId);
            missing.innerHTML = "";
            monikers.sort(function(a, b) {
                const nameA = a.toUpperCase();
                const nameB = b.toUpperCase();
                if (nameA < nameB) {return -1;}
                if (nameA > nameB) {return 1;}
                return 0;
            });
            monikers.forEach((moniker) => {
                let li = document.createElement("li")
                li.appendChild(document.createTextNode(moniker));
                missing.appendChild(li);
            });
        }

        let updating = true;
        let pausedAt = 0;
        let pauseOffset = 0;
        missingChart.on('click', function (e){
            if (e.hasOwnProperty('dataIndex') && e.dataIndex < data.blocks.length) {
                updating = false;
                pausedAt = data.blocks[e.dataIndex-pauseOffset]
                fetch("/block?num="+pausedAt, {
                    method: 'GET',
                    mode: 'cors',
                    cache: 'no-cache',
                    credentials: 'same-origin',
                    redirect: 'error',
                    referrerPolicy: 'no-referrer'
                }).then(event => {
                    event.json().then(upd => {
                        if (upd.hasOwnProperty("missing")) {
                            setMissing("🛑 Missed " + pausedAt + ": ", upd, false)
                        }
                    });
                })
            }
        });
        missingChart.on('globalout', function (){
            if (!updating) {
                const missingWhen = document.getElementById('missingWhen');
                missingWhen.innerHTML = "⏲ Missed " + pausedAt + ":<br/>&nbsp;<br/>&nbsp;";
                updating = true;
            }
        })

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }
        function connectChart() {
            const socket = new WebSocket(wsProto + location.host + '/chart/ws');
            socket.addEventListener('message', function (event) {
                const upd = JSON.parse(event.data);
                data.blocks.shift();
                data.blocks.push(upd.block);
                data.time.shift();
                data.time.push(upd.time);
                data.missed.shift();
                data.missed.push(upd.missed);
                data.missing_percent.shift();
                data.missing_percent.push(upd.missing_percent);
                data.took.shift();
                data.took.push(upd.took);
                if (updating) {
                    missingChart.setOption(option);
                }
            });
            socket.onclose = function(e) {
                console.log('Socket is closed, retrying /chart/ws ...', e.reason);
                setTimeout(function() {
                    connectChart();
                }, 4000);
            };
        }
        connectChart()

        let locked = false;
        function connectMissed() {
            const tableSocket = new WebSocket(wsProto + location.host + '/missed/ws');
            tableSocket.addEventListener('message', function (event) {
                let upd = JSON.parse(event.data);
                if (updating && !locked) {
                    locked = true;
                    setTimeout(function () {
                        if (!updating) {
                            locked = false
                            return
                        }
                        pauseOffset = 0;
                        setMissing("Currently Missing:", upd, false)
                        setMissing("Jailed (unbonding):", upd, true)
                        locked = false
                    }, 1000)
                } else {
                    pauseOffset += 1;
                }
                document.getElementById('headblock').innerHTML = upd.block_num;
                document.getElementById('seconds').innerHTML = upd.delta_sec;
            });
            tableSocket.onclose = function(e) {
                console.log('Socket is closed, retrying /missed/ws ...', e.reason);
                setMissing("⚠ Not Connected", {missing:{"error": ""}}, true)
                setMissing("⚠ Not Connected", {missing:{"error": ""}}, false)
                document.getElementById('headblock').innerHTML = "unknown";
                document.getElementById('seconds').innerHTML = "⚠ ";
                setTimeout(function() {
                    connectMissed();
                }, 4000);
            };
        }
        connectMissed()

    } catch (e) {
        console.log(e.toString());
    }
}
