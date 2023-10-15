import React, { useMemo, useState } from 'react';
import '../../css/map.css';
import { GoogleMap, MarkerF, useLoadScript } from "@react-google-maps/api";

export default function Maps() {

  const { isLoaded } = useLoadScript({
    googleMapsApiKey: process.env.REACT_APP_GOOGLE_MAPS_API_KEY
  })
  const [markers, setMarkers] = useState([]);
  
  if(!isLoaded) return <h1>Loading....</h1>

  return (
    <>
      <Map />
    </>
  )

  function Map() {
    const center = useMemo(() => ({lat: 44, lng: -80}), []);

    const handleMapClick = (event) => {
      const newMarker = {lat: event.latLng.lat(), lng: event.latLng.lng()}
      setMarkers((prev) => [...prev, newMarker]);
    }
    
    return(
      <>
        <GoogleMap zoom={10} 
          center={center} 
          mapContainerClassName='map-container'
          onClick={handleMapClick}
        >
          
          {markers.map((marker, idx) => (
            <MarkerF key={idx} position={marker} draggable/>
          ))}
        </GoogleMap>
      </>
    )
  }
}
