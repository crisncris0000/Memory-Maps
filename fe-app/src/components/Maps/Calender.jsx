import React from 'react';
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers';
import { zhCN } from '@mui/x-date-pickers/locales';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';

export default function Calender() {
    const [selectedDate, setSelectedDate] = React.useState(new Date());

    const handleChange = (date) => {
      setSelectedDate(date);
    };
  
    return (
      <LocalizationProvider dateAdapter={AdapterDateFns} locale={zhCN}>
        <DatePicker
          value={selectedDate}
          onChange={handleChange}
        />
      </LocalizationProvider>
    );
}
