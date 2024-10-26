import { Button } from "@/app/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/app/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { CoworkingDTO } from "@contracts/dtos/coworking";
import Image from "next/image";

interface CoworkingCardProperties {
  coworking: CoworkingDTO;
}

export function CoworkingCard(properties: CoworkingCardProperties) {
  const { coworking } = properties;

  const coworkingImagePath = `http://localhost:4566/${coworking.imagePath[0]}`;

  return (
    <>
      <Card className="group max-w-[320px] space-y-4 border-t-4 p-4">
        <CardHeader>
          <CardTitle>{coworking.name}</CardTitle>
          <CardDescription>Wating for you.</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid w-full items-center gap-4">
            <Image src={coworkingImagePath} width={500} height={500} alt="Picture of the author" />
            <ScrollArea className="h-20 w-60 rounded-md">
              <div className="p-4 ">
                <h4 className="mb-4 text-sm font-medium leading-none">Workspaces:</h4>
                {coworking.workspaces.map((workspace) => (
                  <>
                    <div key={workspace.UUID} className="text-sm flex justify-between items-center">
                      <span>{workspace.name}</span>
                      <span>{workspace.seats}</span>
                    </div>
                    <Separator className="my-2" />
                  </>
                ))}
              </div>
            </ScrollArea>
          </div>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline">Cancel</Button>
          <Button>Deploy</Button>
        </CardFooter>
      </Card>
    </>
  );
}
