import React from 'react';
import Map from 'react-map-gl';
import '../../css/map.css'

export default function Maps() {
  return (
    <div className="map-container">
        <Map
            mapboxAccessToken={process.env.REACT_APP_MAPS_API_KEY}
            initialViewState={{
                longitude: -122.4,
                latitude: 37.8,
                zoom: 10
            }}
            style={{width: 1000, height: 700}}
            mapStyle="mapbox://styles/mapbox/streets-v9"
        />
    </div>
  )
}
