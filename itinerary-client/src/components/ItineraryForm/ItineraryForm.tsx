import React, { ReactHTMLElement, useCallback, useState } from "react";
import axios from "axios";
import Itinerary from "../../types";

interface ItineraryFormProps {
  itinerary: Itinerary;
}

export const ItineraryForm = ({ itinerary }: ItineraryFormProps) => {
  console.log("itineraryform", itinerary.itinerary_name);

  return (
    <>
      <div>
        <div>{itinerary.itinerary_name}</div>
        <div>{itinerary.start_date}</div>
        <div>{itinerary.end_date}</div>
        <div>{itinerary.locations}</div>
      </div>
    </>
  );
};
