import React from 'react'
import Itinerary from '../types'

interface ListProps {
    itineraries: Itinerary[]
}

const Itineraries: Itinerary[] = [
    {
        id: "12345",
        name: "Summer 2022",
        items: ['ice cream', 'mini golf', 'fishing'],
        startDate: "07/01/2022",
        endDate: "07/11/2022",
        locations: ['Stone Harbor, NJ'],
    },
    {
        id: "54321",
        name: "Summer 2023",
        items: ['', '', ''],
        startDate: "08/01/2021",
        endDate: "08/15/2021",
        locations: [],
    }
]

export const ItineraryList = () => {

// export const ItineraryList = ({ itinerary }: ListProps ) => {
    return <div>List goes here</div>
}