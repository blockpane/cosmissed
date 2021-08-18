async function getGeo() {
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

        const response = await fetch("/map", {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer'
        });
        let pex = await response.json()

        const chartDom = document.getElementById('globe');
        const globeChart = echarts.init(chartDom);

        let option;
        option = {
            backgroundColor: 'rgba(7, 7, 7, 1)',
            globe: {
                baseTexture: '/img/world.topo.bathy.200401.jpg',
                heightTexture: '/img/bathymetry_bw_composite_4k.jpg',
                shading: 'color',
                light: {
                    ambient: {
                        intensity: 0.1
                    },
                    main: {
                        intensity: 0.0
                    }
                },
                viewControl: {
                    autoRotate: false,
                    minDistance: 20,
                    //autoRotateSpeed: 4,
                    //distance: 200,
                    //alpha: 25,
                }
            },
            series: {
                type: 'lines3D',
                coordinateSystem: 'globe',
                blendMode: 'lighter',
                lineStyle: {
                    width: 1,
                    color: 'rgb(90, 45, 0)',
                    //color: 'rgb(32,25,73)',
                    opacity: 0.3
                },
                effect: {
                    show: true,
                    trailWidth: 2,
                    trailOpacity: 0.8,
                    trailLength: 0.05,
                    period: 5,
                    //constantSpeed: 40
                },
                data: pex
            }
        };

        option && globeChart.setOption(option);

        const sunDom = document.getElementById('sunburst');
        const sunChart = echarts.init(sunDom);

        let sunOption;
        sunOption = {
            textStyle: {
                color: 'rgba(255, 255, 255, 0.7)'
            },
            backgroundColor: 'rgba(255, 255, 255, 0.0)',
            title: {
                text: 'City in Top 5 Countries',
                left: 'center'
            },
            visualMap: {
                type: 'continuous',
                min: 0,
                max: 10,
                inRange: {
                    //color: ['rgb(89,71,190)', '#537b13', '#ceaf24', 'rgb(232,133,31)']
                    color: ['rgb(94,39,0)', 'rgb(255,138,22)']
                },
            },
            tooltip: {
                trigger: 'item'
            },
            series: {
                type: 'sunburst',
                data: netParams.sunburst.slice(0,5),
                radius: [0, '90%'],
                //label: {
                //    rotate: 'radial'
                //}
                emphasis: {
                    focus: 'ancestor'
                },

                levels: [{},{
                        r0: '0',
                        r: '25%',
                        itemStyle: {
                            borderWidth: 1
                        },
                        label: {
                            rotate: 'tangential'
                        }
                    }, {
                        r0: '25%',
                        r: '75%',
                        label: {
                            align: 'left'
                        }
                    },
                ],
            },

        };
        sunOption && sunChart.setOption(sunOption);

        let wsProto = "ws://"
        if (location.protocol === "https:") {
            wsProto = "wss://"
        }
        function connectGlobe() {
            const socket = new WebSocket(wsProto + location.host + '/map/ws');
            socket.addEventListener('message', function (event) {
                const upd = JSON.parse(event.data);
                option.series.data = upd
                globeChart.setOption(option);
            });
            socket.onclose = function(e) {
                console.log('Socket is closed, retrying /map/ws ...', e.reason);
                setTimeout(function() {
                    connectGlobe();
                }, 4000);
            };
        }
        connectGlobe()

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
                sunOption.series.data = updNet.sunburst.slice(0,5);
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