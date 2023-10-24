import React from 'react'
import Itinerary from '../types'

interface Props {
    itinerary: Itinerary
}

const Itineraries: Itinerary[] = [
    {
        id: "",
        name: "",
        items: [],
        startDate: "",
        endDate: "",
        locations: [],
    }
]

export const ItineraryList = () => {

// export const ItineraryList: React.FC<Props> = ({ itinerary }) => {
    // console.log(itinerary)
    return <div>Hi</div>
}