import React, { useEffect, useRef } from 'react';
import '../../css/map.css'
import mapboxgl from 'mapbox-gl';
import { SearchBox } from '@mapbox/search-js-react'


export default function Maps() {
    const mapContainerRef = useRef(null);
    const mapRef = useRef(null)

    const handleOnRetrieve = (result) => {
       const coords = result.features[0].geometry.coordinates

       const longitude = coords[0]
       const latitude = coords[1]

       console.log(longitude, latitude)

       mapRef.current.flyTo({
            center: [longitude, latitude],
            zoom: 20
       })
    }

    useEffect(() => {

        mapboxgl.accessToken= process.env.REACT_APP_MAPS_API_KEY;

        const map = new mapboxgl.Map({
            container: mapContainerRef.current,
            style: "mapbox://styles/mapbox/streets-v12",
        })
        mapRef.current = map
    }, [])


    return (
        <>
            <div className="search-box-container">
                <SearchBox
                accessToken={process.env.REACT_APP_MAPS_API_KEY} 
                value=''
                onRetrieve={handleOnRetrieve}/>
            </div>
            <div className="map-container">
                <div className="map-content">
                    <div ref={mapContainerRef}
                    style={{width: 1000, height: 700}}></div>
                </div>
            </div>
        </>
    );
}
