import React, { useState } from "react";
import Itinerary from "../../types";
import { Link } from "react-router-dom";
import "../itinerary.css";

interface ListProps {
  itineraries: Itinerary[];
}

export const ItineraryList = ({ itineraries }: ListProps) => {
  // const [itineraryList, setItineraryList] = useState(itineraries);

  const itineraryListItems = itineraries.map((i) => {
    return (
      <div key={i.id}>
        <div>
          <Link to={`/details/${i.id}`}>
                    <h3>Name: {i.itinerary_name}</h3>

          </Link>
        </div>
        <div>Start Date: {i.start_date}</div>
        <div>End Date:{i.end_date}</div>
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
