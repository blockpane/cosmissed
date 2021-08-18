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
        document.getElementById('lastUpdate').innerHTML = netParams.last_updated

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
            backgroundColor: '#000',
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
                    trailWidth: 1,
                    trailOpacity: 0.8,
                    trailLength: 0.05,
                    period: 5,
                    //constantSpeed: 40
                },
                data: pex
            }
        };

        option && globeChart.setOption(option);

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
                document.getElementById('lastUpdate').innerHTML = updNet.last_updated
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