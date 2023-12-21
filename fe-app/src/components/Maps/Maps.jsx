import React, { useEffect, useRef, useState } from 'react';
import '../../css/map.css';
import mapboxgl from 'mapbox-gl';
import { SearchBox } from '@mapbox/search-js-react';
import LocationInfo from './LocationInfo';
import axios from 'axios';
import MarkerInfo from './MarkerInfo';
import Navigation from './Navigation';
import Calender from './Calender';


export default function Maps() {
    const mapContainerRef = useRef(null);
    const mapRef = useRef(null)
    const markerRef = useRef(null);

    const [showForm, setShowForm] = useState(false);
    const [showMarkerInfo, setShowMarkerInfo] = useState(false);
    const [selectedMarker, setSelectedMarker] = useState(null);
    const [longitude, setLongitude] = useState(null);
    const [latitude, setLatitude] = useState(null);
    const [markerPosts, setMarkerPosts] = useState(null);

    const handleOnRetrieve = (result) => {
       const coords = result.features[0].geometry.coordinates;

       mapRef.current.flyTo({
        center: [coords[0], coords[1]],
        zoom: 20
    });
        
       setLongitude(coords[0]);
       setLatitude(coords[1]);
       setShowForm(true);
    }

    const handleOnClose = () => {
        if(markerRef.current) {
            markerRef.current.remove()
        }
    }

    const getMarkerPosts = () => {
        markerPosts.forEach((post) => {
            const marker = new mapboxgl.Marker()
            .setLngLat([post.longitude, post.latitude])
            .addTo(mapRef.current);

            marker.getElement().addEventListener('click', () => {
                setShowMarkerInfo(true);
                setSelectedMarker(post);
            })
            
            markerRef.current = marker;
        })
    }

    useEffect(() => {

        axios.get("http://localhost:8080/marker-posts").then((response) => {
            setMarkerPosts(response.data.markerposts);
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
            {markerPosts ? getMarkerPosts() : null}
            <Navigation />
            <Calender setMarkerPosts={setMarkerPosts} markerPosts={markerPosts}/>
            <LocationInfo show={showForm} setShow={setShowForm} longitude={longitude} latitude={latitude} onHide={handleOnClose}/>
            <MarkerInfo show={showMarkerInfo} setShow={setShowMarkerInfo} markerPost={selectedMarker}/>
            <div className="map-container">
                <div className="map-content">
                    <div ref={mapContainerRef}
                    style={{width: "100%", height: "100%", borderRadius: "5px"}}>
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