import React, { useState } from 'react';
import Map, { Marker, Popup} from 'react-map-gl';
import { AddressAutofill } from '@mapbox/search-js-react';
import '../../css/map.css';
import LocationInfo from './LocationInfo';

export default function Maps() {
  const [markers, setMarkers] = useState([]);
  const [showPopup, setShowPopup] = useState(false);
  const [show, setShow] = useState(false);
  const [value, setValue] = useState('');

  const addMarker = (event) => {
    const longitude = event.lngLat.lng;
    const latitude = event.lngLat.lat;

    setShow(true);    

    //setMarkers([...markers, { longitude, latitude }]);
  };

  return (
    <>
      <form>
        <AddressAutofill accessToken={process.env.REACT_APP_MAPS_API_KEY}>
          <input
          autoComplete="shipping address-line1"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          />
        </AddressAutofill>
      </form>
      <Map
        mapboxAccessToken={process.env.REACT_APP_MAPS_API_KEY}
        initialViewState={{
          longitude: -100,
          latitude: 40,
          zoom: 3.5
        }}
        mapStyle="mapbox://styles/mapbox/streets-v9"
        style={{ width: "100vw", height: "100vh" }}
        onClick={addMarker}
      >
        
        <LocationInfo show={show} setShow={setShow}/>

        {markers.map((marker, index) => (
          <Marker key={index} longitude={marker.longitude} 
          latitude={marker.latitude} offsetTop={-20} offsetLeft={-10} />
        ))}
      </Map>
    </>
  );
}