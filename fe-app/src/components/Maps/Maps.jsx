import React from 'react';
import Map, { Marker } from 'react-map-gl';
import '../../css/map.css';

export default function Maps() {
  return (
      <Map
          mapboxAccessToken={process.env.REACT_APP_MAPS_API_KEY}
          initialViewState={{
            longitude: -100,
            latitude: 40,
            zoom: 3.5
          }}
          mapStyle="mapbox://styles/mapbox/streets-v9"
          style={{width: "100vw", height: "100vh"}}
      >
          <Marker longitude={-100} latitude={40} offsetTop={-20} offsetLeft={-10}> 
          </Marker>
      </Map>
  )
}
