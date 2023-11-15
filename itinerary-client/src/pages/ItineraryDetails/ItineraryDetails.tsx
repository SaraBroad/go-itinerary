import React, { useState } from "react";
import { ItineraryDetails } from "../../components/ItineraryDetails/ItineraryDetails";

export const ItineraryDetailsPage = () => {
    const [isCompleted, setIsCompleted] = useState(false)
    
    return (
        <div>
            <div>
                <ItineraryDetails 
                    name={"The Louvre"} 
                    isCompleted={true} 
                />
            </div>

        </div>
    )
}