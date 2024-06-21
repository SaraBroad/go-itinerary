import React, { useState } from "react";
import { ItineraryDetails } from "../../components/ItineraryDetails/ItineraryDetails";
import { useParams } from "react-router-dom";

export const ItineraryDetailsPage = () => {
  const [isCompleted, setIsCompleted] = useState(false);
  const { id } = useParams()
  console.log("id", id)

  return (
    <div>
      <div>
        <ItineraryDetails name={"The Louvre"} isCompleted={true} />
      </div>
    </div>
  );
};
