import { CoworkingCard } from "@/app/coworkings/components/coworking-card";
import { CoworkingResponse } from "@/core/coworking/responses/coworking.response";
import { CoworkingDTO } from "@contracts/dtos/workspace";

export interface CoworkingGridProperties {
  coworkings: CoworkingDTO[];
}

export function CoworkingCardGrid(properties: CoworkingGridProperties) {
  const { coworkings } = properties;
  return (
    <div className="grid grid-cols-1 place-items-center gap-6 md:grid-cols-2 lg:grid-cols-4">
      {coworkings.map((coworking) => (
        <div key={coworking.UUID}>
          <CoworkingCard coworking={coworking} />
        </div>
      ))}
    </div>
  );
}
