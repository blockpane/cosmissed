async function chainInfo() {
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
        document.getElementById('networkId').innerHTML = params.chain

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }

        function connectMissed() {
            const tableSocket = new WebSocket(wsProto + location.host + '/missed/ws');
            tableSocket.addEventListener('message', function (event) {
                let upd = JSON.parse(event.data);
                document.getElementById('headblock').innerHTML = upd.block_num;
                document.getElementById('missingNow').innerHTML = upd.vote_missing
                document.getElementById('seconds').innerHTML = upd.delta_sec;
                document.getElementById('proposedBy').innerHTML = upd.proposer
            });
            tableSocket.onclose = function(e) {
                console.log('Socket is closed, retrying /missed/ws ...', e.reason);
                document.getElementById('headblock').innerHTML = "⚠"
                document.getElementById('missingNow').innerHTML = "⚠ Not Connected"
                document.getElementById('proposedBy').innerHTML = "⚠"
                document.getElementById('seconds').innerHTML = "⚠ ";
                setTimeout(function() {
                    connectMissed();
                }, 4000);
            };
        }
        connectMissed()

        const netInfo = await fetch("/net", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        const netParams = await netInfo.json()
        document.getElementById('totalNodes').innerHTML = netParams.peers_discovered
        document.getElementById('rpcNodes').innerHTML = netParams.rpc_discovered
        let lupd = netParams.last_updated
        if (lupd === '0001-01-01T00:00:00Z') {
            lupd = 'Please standby, updating peers can take up to 5 minutes.'
        }
        document.getElementById('lastUpdate').innerHTML = lupd

        const locDom = document.getElementById('cities');
        const locChart = echarts.init(locDom);

        let locOption;
        locOption = {
            textStyle: {
                color: 'rgba(255, 255, 255, 0.7)'
            },
            backgroundColor: 'rgba(255, 255, 255, 0.0)',
            title: {
                text: 'Top 10 Node Countries',
                left: 'center'
            },
            visualMap: {
                type: 'continuous',
                min: 0,
                max: 50,
                inRange: {
                    //color: ['rgb(89,71,190)', '#537b13', '#ceaf24', 'rgb(232,133,31)']
                    color: ['rgb(135,87,2)', 'rgb(255,162,0)']
                },
            },
            tooltip: {
                trigger: 'item'
            },
            series: {
                type: 'sunburst',
                data: netParams.sunburst.slice(0, 100),
                radius: [0, '90%'],
                emphasis: {
                    focus: 'ancestor'
                },
                levels: [{}, {
                    r0: '15',
                    r: '50%',
                    itemStyle: {
                        borderWidth: 1
                    },
                    label: {
                        align: 'right'
                    }
                }, {
                        r0: '50%',
                        r: '55%',
                        label: {
                            position: 'outside',
                            padding: 3,
                            silent: false,
                            color: "#fff",
                        },
                    },
                ],
            },
        };
        locOption && locChart.setOption(locOption);

        const sunDom = document.getElementById('sunburst');
        const sunChart = echarts.init(sunDom);

        let sunOption;
        sunOption = {
            textStyle: {
                color: 'rgba(255, 255, 255, 0.7)'
            },
            backgroundColor: 'rgba(255, 255, 255, 0.0)',
            title: {
                text: 'Top 10 Hosting Providers',
                left: 'center'
            },
            visualMap: {
                type: 'continuous',
                min: 0,
                max: 50,
                inRange: {
                    //color: ['rgb(89,71,190)', '#537b13', '#ceaf24', 'rgb(232,133,31)']
                    color: ['rgb(18,10,31)', 'rgb(107,59,177)']
                },
            },
            tooltip: {
                trigger: 'item'
            },
            series: {
                type: 'sunburst',
                data: netParams.providers.slice(0, 100),
                radius: [0, '90%'],
                emphasis: {
                    focus: 'ancestor'
                },
                levels: [{}, {
                    r0: '10',
                    r: '50%',
                    itemStyle: {
                        borderWidth: 1
                    },
                    label: {
                        align: 'right'
                    }
                }, {
                    r0: '50%',
                    r: '60%',
                    label: {
                        rotate: 'tangential'
                    }
                },
                    {
                        r0: '60%',
                        r: '65%',
                        label: {
                            position: 'outside',
                            padding: 3,
                            silent: false,
                            color: "#fff",
                        },
                    },
                ],
            },
        };
        sunOption && sunChart.setOption(sunOption);

        function connectNet() {
            const socket = new WebSocket(wsProto + location.host + '/net/ws');
            socket.addEventListener('message', function (event) {
                const updNet = JSON.parse(event.data);
                document.getElementById('totalNodes').innerHTML = updNet.peers_discovered
                document.getElementById('rpcNodes').innerHTML = updNet.rpc_discovered
                let lupd = updNet.last_updated
                if (lupd === '0001-01-01T00:00:00Z') {
                    lupd = 'Please standby, updating peers can take up to 5 minutes.'
                }
                document.getElementById('lastUpdate').innerHTML = lupd
                sunOption.series.data = updNet.providers.slice(0,10);
                sunChart.setOption(sunOption);
            });
            socket.onclose = function(e) {
                console.log('Socket is closed, retrying /net/ws ...', e.reason);
                setTimeout(function() {
                    connectNet();
                }, 4000);
            };
        }
        connectNet()

    }
    catch (e) {
        console.log(e)
    }
}