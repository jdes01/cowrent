'use server'

import { CoworkingDTO } from "@contracts/dtos/coworking";
import { unstable_cache as cache } from 'next/cache'

export async function findCoworkings(): Promise<CoworkingDTO[]> {
    const getCachedCoworkings = cache(
        async () => {
            try {

                const response = await fetch("http://localhost:8080/api/coworkings");
                const coworkingsJson = await response.json();

                const coworkings: CoworkingDTO[] = coworkingsJson["coworkings"].map((coworking: any) => {
                    return CoworkingDTO.fromJSON(coworking)
                });

                console.log(coworkings[0])

                return coworkings
            } catch {
                return []
            }
        },
        ['find-coworkings'],
        {
            tags: ['coworkings'],
        },
    )

    return getCachedCoworkings()
}
