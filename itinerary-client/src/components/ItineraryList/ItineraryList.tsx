import React, { useState } from "react";
import Itinerary from "../../types";
import "../itinerary.css";

interface ListProps {
  // itineraries: Itinerary[];
  itineraries: Itinerary;

}

export const ItineraryList = ({ itineraries }: ListProps) => {
  console.log("itineraries props", itineraries);
  const [itineraryList, setItineraryList] = useState(itineraries);

  // const itineraryListItems = itineraries.map((i) => {
  //   return (
  //     <div key={i.id}>
  //       <p>hi</p>
  //       <div>
  //         <h3>Name: {i.itinerary_name}</h3>{" "}
  //       </div>
  //       <div>Start Date:{i.startDate}</div>
  //       <div>End Date:{i.endDate}</div>
  //       {i.items?.length === 0 ? (
  //         <div>Activities: Nothing to do yet</div>
  //       ) : (
  //         <div>Activities:{i.items}</div>
  //       )}
  //       {i.locations?.length === 0 ? (
  //         <div>Locations: No destinations yet</div>
  //       ) : (
  //         <div>Locations:{i.locations}</div>
  //       )}
  //     </div>
  //   );
  // });

  return (
    <div>
      <div className="list-component">
        <div>{itineraries.endDate}</div>
        <div>{itineraries.itinerary_name}</div>
      </div>
    </div>
  );
};
