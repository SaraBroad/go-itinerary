import React from "react";
import { Routes, Route } from "react-router-dom";

import logo from "./logo.svg";
import "./App.css";
import { ItineraryList } from "./components/ItineraryList/ItineraryList";
import { ItineraryFormPage } from "./pages/ItineraryForm/ItineraryForm";
import { ItineraryListPage } from "./pages/ItineraryList/itineraryList";
import { Header } from "./components/Header/Header";
import Itinerary from "./types";

const itineraries: Itinerary[] = [
  {
    id: "12345",
    name: "Summer 2022",
    items: ["ice cream", "mini golf", "fishing"],
    startDate: "07/01/2022",
    endDate: "07/11/2022",
    locations: ["Stone Harbor, NJ"],
  },
  {
    id: "54321",
    name: "Summer 2023",
    items: ["Berlin Wall ", "Charles Bridge ", "Currywurst"],
    startDate: "08/01/2021",
    endDate: "08/15/2021",
    locations: ["Berlin", "Prague"],
  },
  {
    id: "932923",
    name: "Summer 2024",
    items: [],
    startDate: "06/15/2024",
    endDate: "06/27/2024",
    locations: [],
  },
];

function App() {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/list" element={<ItineraryListPage />} />
        <Route path="/newItinerary" element={<ItineraryFormPage />} />
        {/* <Route path="/list" element={<ItineraryListPage itineraries={itineraries}/>} /> */}
        {/* <Route path="/list" element={<ItineraryList/>} /> */}
      </Routes>
    </>
  );
}

export default App;
