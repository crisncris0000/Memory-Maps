import React, { useState } from 'react';
import Map, { GeolocateControl, Marker} from 'react-map-gl';
import '../../css/map.css';
import LocationInfo from './LocationInfo';
import SearchBar from './SearchBar';

export default function Maps() {
  const [markers, setMarkers] = useState([]);
  const [show, setShow] = useState(false);
  const [viewState, setViewState] = useState({
    longitude: -100,
    latitude: 40,
    zoom: 3.5
  })

  const showForm = (event) => {
    const longitude = event.lngLat.lng;
    const latitude = event.lngLat.lat;

    setShow(true);    

    //setMarkers([...markers, { longitude, latitude }]);
  };

  const handleGeolocation = (position) => {
    const {longitude, latitude} = position.coords;
    setMarkers([...markers, { longitude, latitude }]);
  }

  return (
    <>
      <SearchBar markers={markers} setMarkers={setMarkers} viewState={viewState} setViewState={setViewState}/>
      <div className="map-container">
        <div className="map-content">
          <Map
            mapboxAccessToken={process.env.REACT_APP_MAPS_API_KEY}
            {...viewState}
            mapStyle="mapbox://styles/mapbox/streets-v9"
            style={{ width: 1000, height: 700 }}
            onClick={showForm}
          >

            
            <LocationInfo show={show} setShow={setShow}/>

            {markers.map((marker, index) => (
              <Marker key={index} longitude={marker.longitude} 
              latitude={marker.latitude} offsetTop={-20} offsetLeft={-10} />
            ))}

            <GeolocateControl
            position="top-left"
            trackUserLocation
            onGeolocate={handleGeolocation}
            />
          </Map>
        </div>
      </div>
    </>
  );
}