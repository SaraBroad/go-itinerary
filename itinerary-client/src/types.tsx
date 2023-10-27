export default interface Itinerary {
    id: string
    name: string
    items?: string[]
    startDate?: string
    endDate?: string
    locations?: string[]
}