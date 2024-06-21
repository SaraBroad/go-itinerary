import React, { useState } from "react";
import { ItineraryDetails } from "../../components/ItineraryDetails/ItineraryDetails";
import { useParams } from "react-router-dom";

export const ItineraryDetailsPage = () => {
  const [isCompleted, setIsCompleted] = useState(false);
  const [locations, setLocations] = useState([])
  const [items, setItems] = useState([])
  const { id } = useParams()
  console.log("id", id)

  const handleLocationSubmit = (e: { preventDefault: () => void }) => {
    alert(JSON.stringify(locations, null, 2));
  }
  
  const handleItemsSubmit = (e: { preventDefault: () => void }) => {
    alert(JSON.stringify(items, null, 2));
  }

  return (
    <div>
      <div>
        <ItineraryDetails name={"The Louvre"} isCompleted={true} />
        <form className="form" onSubmit={handleLocationSubmit}>

        </form>
      </div>
    </div>
  );
};
