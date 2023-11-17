import React, { useEffect, useRef, useState } from 'react';
import '../../css/map.css';
import mapboxgl from 'mapbox-gl';
import { SearchBox } from '@mapbox/search-js-react';
import LocationInfo from './LocationInfo';
import axios from 'axios';


export default function Maps() {
    const mapContainerRef = useRef(null);
    const mapRef = useRef(null)
    const markerRef = useRef(null);

    const [show, setShow] = useState(false);
    const [longitude, setLongitude] = useState(null);
    const [latitude, setLatitude] = useState(null);

    const handleOnRetrieve = (result) => {
       const coords = result.features[0].geometry.coordinates;

       setLongitude(coords[0]);
       setLatitude(coords[1]);
       setShow(true);

       mapRef.current.flyTo({
            center: [longitude, latitude],
            zoom: 20
       });
    }

    const handleOnClose = () => {
        if(markerRef.current) {
            markerRef.current.remove()
        }
    }

    useEffect(() => {

        axios.get("http://localhost:8080/marker-posts").then((response) => {
            response.data.markerposts.forEach((post) => {
                const marker = new mapboxgl.Marker()
                .setLngLat([post.longitude, post.latitude])
                .addTo(mapRef.current);
                
                markerRef.current = marker;
            })
        }).catch((error) => {
            console.log(error);
        });


        mapboxgl.accessToken= process.env.REACT_APP_MAPS_API_KEY;

        const map = new mapboxgl.Map({
            container: mapContainerRef.current,
            style: "mapbox://styles/mapbox/streets-v12",
        });

        mapRef.current = map;
        
    }, [])


    return (
        <>
            <LocationInfo show={show} setShow={setShow} longitude={longitude} latitude={latitude} onHide={handleOnClose}/>
            <div className="map-container">
                <div className="map-content">
                    <div ref={mapContainerRef}
                    style={{width: "100%", height: "100%"}}>
                        <div className="search-box-container">
                            <SearchBox
                            accessToken={process.env.REACT_APP_MAPS_API_KEY} 
                            value=''
                            onRetrieve={handleOnRetrieve}/>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
