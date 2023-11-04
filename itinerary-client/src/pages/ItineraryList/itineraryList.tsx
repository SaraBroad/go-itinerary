import React from "react";
import Itinerary from "../../types";
import { ItineraryList } from "../../components/ItineraryList/ItineraryList";
  
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

  export const ItineraryListPage = ({ itineraries }: {itineraries: Itinerary[] }) => {
    console.log(itineraries)
    return (
        <>
        <h3 className="itinerary-header">Itineraries</h3>
        <ItineraryList itineraries={itineraries} />
        {/* <ItineraryList /> */}
        </>
    )
}