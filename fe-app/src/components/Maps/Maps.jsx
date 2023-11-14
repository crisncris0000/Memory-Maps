import React, { useEffect, useRef, useState } from 'react';
import '../../css/map.css';
import mapboxgl from 'mapbox-gl';
import { SearchBox } from '@mapbox/search-js-react';
import LocationInfo from './LocationInfo';


export default function Maps() {
    const mapContainerRef = useRef(null);
    const mapRef = useRef(null)
    const markerRef = useRef(null);

    const [markers, setMarkers] = useState([]);
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
       
       const marker = new mapboxgl.Marker()
       .setLngLat([longitude, latitude])
       .addTo(mapRef.current);
       
       markerRef.current = marker;
    }

    useEffect(() => {

        mapboxgl.accessToken= process.env.REACT_APP_MAPS_API_KEY;

        const map = new mapboxgl.Map({
            container: mapContainerRef.current,
            style: "mapbox://styles/mapbox/streets-v12",
        });

        mapRef.current = map;
        
    }, [])


    return (
        <>
            <LocationInfo show={show} setShow={setShow} longitude={longitude} latitude={latitude}/>
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
