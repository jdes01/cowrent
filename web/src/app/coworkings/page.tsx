import { CoworkingCardGrid } from "@/app/coworkings/components/coworkings-card-grid";
import { findCoworkings } from "@/core/coworking/actions/find-coworkings";

export default async function Page() {
  const coworkings = await findCoworkings();

  return (
    <main>
      <CoworkingCardGrid coworkings={coworkings} />
    </main>
  );
}
