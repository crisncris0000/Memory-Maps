import React from 'react';
import { Map } from 'react-map-gl';
import '../../css/map.css'
import { SearchBox } from '@mapbox/search-js-react'


export default function Maps() {
    return (
        <>
            <div className="search-box-container">
                <SearchBox accessToken={process.env.REACT_APP_MAPS_API_KEY} 
                value=''/>
            </div>
            <div className="map-container">
                <div className="map-content">
                    <Map mapboxAccessToken={process.env.REACT_APP_MAPS_API_KEY}
                        mapStyle="mapbox://styles/mapbox/streets-v9"
                        style={{width: 1000, height: 700}}>
                    </Map>
                </div>
            </div>
        </>
    );
}
