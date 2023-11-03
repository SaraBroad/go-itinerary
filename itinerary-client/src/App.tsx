import React from 'react';
import { Routes, Route } from 'react-router-dom';

import logo from './logo.svg';
import './App.css';
import { ItineraryList } from './components/itinerary-list'
import { ItineraryForm } from './components/itinerary-form';

function App() {
  return (
    <>
    {/* <Routes>
      <Route path="/list" element={<ItineraryList/>} />
    </Routes> */}
      <ItineraryForm /> 
      <ItineraryList />
    </>
  );
}

export default App;
