import React, { useEffect, useState } from "react";
import Itinerary from "../../types";
import { ItineraryList } from "../../components/ItineraryList/ItineraryList";
import { itineraryData } from "../../data";
import axios from "axios";

export const ItineraryListPage = () => {
  // export const ItineraryListPage = ({ itineraries }: {itineraries: Itinerary[] }) => {
  const [itineraryList, setItineraryList] = useState([])  
  // useEffect(() => {
  //   axios.get('')
  // })


  console.log(JSON.stringify(itineraryData));
  return (
    <>
      <h3 className="itinerary-header">Itineraries</h3>
      <ItineraryList itineraries={itineraryData} />
      {/* <ItineraryList /> */}
    </>
  );
};
