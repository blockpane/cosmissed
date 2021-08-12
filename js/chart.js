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
            backgroundColor: '#fff',
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
                    start: 90,
                    end: 100,
                    zoomOnMouseWheel: false,
                    moveOnMouseWheel: false,
                    moveOnMouseMove: false,
                    xAxisIndex: [0, 1],
                    backgroundColor: 'rgba(104,104,104,0.22)',
                    dataBackground: {
                        lineStyle: {
                            color: '#000',
                            width: 2,
                        }
                    }
                },
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
                        color: '#000',
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
                            offset: 1,
                            color: 'rgb(89,71,190)'
                        }])
                    },
                },
                {
                    name: 'missing consensus %',
                    type: 'line',
                    z: 1,
                    step: 'end',

                    // little hack that makes the chart clickable, only the "symbol" will emit a mouse click event,
                    // this will create a giant clear svg and overlay so the whole series is clickable:
                    symbol: 'image://data:image/svg;base64,PD94bWwgdmVyc2lvbj0iMS4wIj8+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjYwMCI+CiAgICA8cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNGRkYiIHN0cm9rZS13aWR0aD0iMCIgZD0ibTAsMGg0ODB2MjcwSDB6Ii8+Cjwvc3ZnPg==',
                    symbolSize: [200, 600],
                    symbolOffset: [0, "50%"],
                    showSymbol: false,

                    hoverAnimation: true,
                    data: data.missing_percent,
                    lineStyle: {
                        color: 'rgb(238,131,25)',
                        width: 2.0,
                        type: 'dashed',
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
                    symbolOffset: [0, "50%"],
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
                        }])
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

        const setMissing = function (title, data){
            let monikers = [];
            for (const [key, value] of Object.entries(data.missing)) {
                monikers.push(key)
            }
            const missingWhen = document.getElementById('missingWhen');
            missingWhen.innerHTML = title
            const missing = document.getElementById('missing');
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
                            setMissing("⏸ Block " + pausedAt + ": ", upd)
                        }
                    });
                })
            }
        });
        missingChart.on('globalout', function (){
            if (!updating) {
                const missingWhen = document.getElementById('missingWhen');
                missingWhen.innerHTML = "▶️ Block " + pausedAt + ":";
                updating = true;
            }
        })

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }
        const socket = new WebSocket(wsProto+location.host+'/chart/ws');
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

        let locked = false;
        const tableSocket = new WebSocket(wsProto+location.host+'/missed/ws');
        tableSocket.addEventListener('message', function (event) {
            let upd = JSON.parse(event.data);
            if (updating && !locked) {
                locked = true;
                setTimeout(function (){
                    if (!updating) {
                        locked = false
                        return
                    }
                    pauseOffset = 0;
                    setMissing("Currently Missing:", upd)
                    locked = false
                }, 3000)
            } else {
                pauseOffset += 1;
            }
            document.getElementById('headblock').innerHTML = upd.block_num;
            document.getElementById('seconds').innerHTML = upd.delta_sec;
        });

    } catch (e) {
        console.log(e.toString());
    }
}
