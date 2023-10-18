import React, { useState } from 'react';
import axios from 'axios';

export default function SearchBar({ markers, setMarkers }) {

    const [suggestions, setSuggestions] = useState([]);
    const [selectedAddress, setSelectedAddress] = useState(null);

    const fetchAddressSuggestions = async (query) => {
        const API_URL = `https://api.mapbox.com/geocoding/v5/mapbox.places/${encodeURIComponent(query)}.json?access_token=${process.env.REACT_APP_MAPS_API_KEY}`;

        try {
            const response = await axios.get(API_URL);
            return response.data.features;
        } catch (error) {
            console.log(error);
            return [];
        }
    }

    const handleAddressChange = async (event) => {
        const query = event.target.value;

        if(query) {
            const results = await fetchAddressSuggestions(query);
            setSuggestions(results);
        } else {
            setSuggestions([]);
        }
    }

    const handleSelectChange = (event) => {
        const selected = suggestions.find(suggestion => suggestion.id === event.target.value);
        setSelectedAddress(selected.place_name);
        if (selected) {
            const [longitude, latitude] = selected.center;
            setMarkers([...markers, { latitude, longitude }]);
        }

        console.log(selected);
    };

    return (
        <div className="search-container">
            <input 
                type="text" 
                onChange={handleAddressChange} 
                placeholder="Search for an address..." 
                className="search-input"
            />
            {suggestions.length > 0 && (
                <select 
                    onChange={handleSelectChange} 
                    value={selectedAddress}
                    className="address-dropdown"
                >
                    <option value="" disabled>Select an address</option>
                    {suggestions.map((feature) => (
                        <option key={feature.id} value={feature.id}>
                            {feature.place_name}
                        </option>
                    ))}
                </select>
            )}
        </div>
    );
}
