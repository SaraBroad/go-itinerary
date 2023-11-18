import React, { useState } from "react";
import Itinerary from "../../types";
import "../itinerary.css";

interface ListProps {
  itineraries: Itinerary[];
}

export const ItineraryList = ({ itineraries }: ListProps) => {
  console.log("itineraries", itineraries);
  const [itineraryList, setItineraryList] = useState(itineraries);

  const itineraryListItems = itineraryList.map((i) => {
    return (
      <div key={i.id}>
        <div>
          <h3>Name: {i.name}</h3>{" "}
        </div>
        <div>Start Date:{i.startDate}</div>
        <div>End Date:{i.endDate}</div>
        {i.items?.length === 0 ? (
          <div>Activities: Nothing to do yet</div>
        ) : (
          <div>Activities:{i.items}</div>
        )}
        {i.locations?.length === 0 ? (
          <div>Locations: No destinations yet</div>
        ) : (
          <div>Locations:{i.locations}</div>
        )}
      </div>
    );
  });

  return (
    <div>
      <div className="list-component">
        <div>{itineraryListItems}</div>
      </div>
    </div>
  );
};
