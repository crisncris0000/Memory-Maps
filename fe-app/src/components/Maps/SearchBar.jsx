import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Select from 'react-select';

export default function SearchBar({ markers, setMarkers }) {

    const [suggestions, setSuggestions] = useState([]);
    const [selectedAddress, setSelectedAddress] = useState(null);
    const [inputValue, setInputValue] = useState('');


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

    const handleInputChange = (inputValue) => {
        setInputValue(inputValue);
        return inputValue;
    }

    const handleChange = option => {
        setSelectedAddress(option.label);
        const [longitude, latitude] = option.center;
        setMarkers([...markers, { latitude, longitude }]);
    }

    useEffect(() => {
        const fetchSuggestions = async () => {
            if (inputValue) {
                const results = await fetchAddressSuggestions(inputValue);
                setSuggestions(results.map(feature => ({
                    value: feature.id,
                    label: feature.place_name,
                    center: feature.center
                })));
            } else {
                setSuggestions([]);
            }
        }
        fetchSuggestions();
    }, [inputValue]);
    

    return (
        <div className="search-container">
            <Select
                onInputChange={handleInputChange}
                onChange={handleChange}
                options={suggestions}
                value={selectedAddress ? { label: selectedAddress } : null}
                placeholder="Search for an address..."
                className="address-dropdown"
                noOptionsMessage={() => "No results found"}
                isLoading={suggestions.length === 0}
                filterOption={false}
            />
        </div>
    );
}
