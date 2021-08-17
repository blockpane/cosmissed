async function getGeo() {
    try {
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
                        intensity: 0.2
                    },
                    main: {
                        intensity: 0.00
                    }
                },
                viewControl: {
                    autoRotate: true,
                    autoRotateSpeed: 4,
                    distance: 200,
                    alpha: 25,
                }
            },
            series: {
                type: 'lines3D',
                coordinateSystem: 'globe',
                blendMode: 'lighter',
                lineStyle: {
                    width: 3,
                    color: 'rgb(90, 45, 0)',
                    opacity: 0.8
                },
                effect: {
                    show: true,
                    trailWidth: 4,
                    trailOpacity: 0.3,
                    trailLength: 0.4,
                    constantSpeed: 100
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
    }
    catch (e) {
        console.log(e)
    }
}