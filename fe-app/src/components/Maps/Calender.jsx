import React from 'react';
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers';
import { zhCN } from '@mui/x-date-pickers/locales';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';

export default function Calender({ setMarkerPosts, markerPosts }) {
    const [selectedStartDate, setSelectedStartDate] = React.useState(new Date());
    const [selectedEndDate, setSelectedEndDate] = React.useState(new Date());

    const handleStartDate = (date) => {
      setSelectedStartDate(date);
    }

    const handleEndDate = (date) => {
      setSelectedEndDate(date);
    }
  
    return (
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
    );
}
