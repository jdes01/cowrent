import { CoworkingCard } from "@/app/coworkings/components/coworking-card";
import { CoworkingCardV0 } from "@/components/coworking-card-v0";
import { CoworkingDTO } from "@contracts/dtos/coworking";

export interface CoworkingGridProperties {
  coworkings: CoworkingDTO[];
}

export function CoworkingCardGrid(properties: CoworkingGridProperties) {
  const { coworkings } = properties;

  return (
    <div className="grid grid-cols-1 place-items-center gap-6 md:grid-cols-2 lg:grid-cols-4">
      {coworkings
        .filter((coworking) => coworking.imagePath.length > 0)
        .map((coworking) => (
          <div key={coworking.UUID}>
            {/* <CoworkingCard coworking={coworking} /> */}
            <CoworkingCardV0 coworking={coworking} />
          </div>
        ))}
    </div>
  );
}
