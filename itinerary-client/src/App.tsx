import React from 'react';
import logo from './logo.svg';
import './App.css';
import { ItineraryList } from './components/itinerary-list'
import { ItineraryForm } from './components/itinerary-form';

function App() {
  return (
    <>
      <ItineraryForm />
      <ItineraryList />
    </>
  );
}

export default App;
