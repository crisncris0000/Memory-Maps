import React, { useState, useEffect } from 'react';
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers';
import { zhCN } from '@mui/x-date-pickers/locales';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';

export default function Calender({ setMarkerPosts, markerPosts }) {
    const [selectedStartDate, setSelectedStartDate] = useState(new Date());
    const [selectedEndDate, setSelectedEndDate] = useState(new Date());
    const [originalMarkerPosts, setOriginalMarkerPosts] = useState([]);

    useEffect(() => {
        setOriginalMarkerPosts(markerPosts);
    }, [originalMarkerPosts]);

    const handleStartDate = (date) => {
        setSelectedStartDate(date);
    }

    const handleEndDate = (date) => {
        setSelectedEndDate(date);
    };

    const handleSubmit = (e) => {
        e.preventDefault();

        const startDateString = new Date(selectedStartDate).toLocaleDateString();
        const endDateString = new Date(selectedEndDate).toLocaleDateString();

        const filteredArray = originalMarkerPosts.filter((post) => {
            const postDateString = new Date(post.createdAt).toLocaleDateString();
            return postDateString >= startDateString && postDateString <= endDateString;
        });

        setMarkerPosts(filteredArray);

    };

    return (
        <form onSubmit={handleSubmit}>
            <LocalizationProvider dateAdapter={AdapterDateFns} locale={zhCN}>
                <DatePicker
                    value={selectedStartDate}
                    onChange={handleStartDate}
                />

                <DatePicker
                    value={selectedEndDate}
                    onChange={handleEndDate}
                    minDate={selectedStartDate}
                />
            </LocalizationProvider>

            <button type="submit">Filter</button>
        </form>
    );
}
