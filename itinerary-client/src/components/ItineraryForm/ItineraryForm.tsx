import React, { ReactHTMLElement, useCallback, useState } from "react";
import axios from "axios";
import Itinerary from "../../types";

interface ItineraryFormProps {
  itinerary: Itinerary;
}

export const ItineraryForm = ({ itinerary }: ItineraryFormProps) => {
  console.log("itineraryform", itinerary.name);

  return (
    <>
      <div>
        <div>{itinerary.name}</div>
        <div>{itinerary.startDate}</div>
        <div>{itinerary.endDate}</div>
        <div>{itinerary.locations}</div>
      </div>
    </>
  );
};
