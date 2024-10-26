"use client";
/**
 * This code was generated by v0 by Vercel.
 * @see https://v0.dev/t/V6I0dFdaqUG
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { JSX, SVGProps, useEffect, useState } from "react";
import { CarouselItem, CarouselContent, CarouselPrevious, CarouselNext, Carousel } from "@/components/ui/carousel";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { CoworkingDTO } from "@contracts/dtos/coworking";

interface CoworkingCardProperties {
  coworking: CoworkingDTO;
}

export function CoworkingCardV0(properties: CoworkingCardProperties) {
  const { coworking } = properties;
  const [currentWorkspaceIndex, setCurrentWorkspaceIndex] = useState(0);

  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurrentWorkspaceIndex((prevIndex) => (prevIndex + 1) % coworking.workspaces.length);
    }, 3000);

    return () => clearInterval(intervalId);
  }, [coworking.workspaces.length]);

  return (
    <Card className="group w-[320px] rounded-lg overflow-hidden shadow-lg">
      <div className="relative">
        <Carousel className="w-full">
          <CarouselContent>
            {coworking.imagePath.map((image, index) => (
              <CarouselItem key={index}>
                <img
                  alt="Event image"
                  className="w-full h-56 object-cover"
                  height="200"
                  src={`http://localhost:4566/${image}`}
                  style={{
                    aspectRatio: "350/200",
                    objectFit: "cover",
                  }}
                  width="350"
                />
              </CarouselItem>
            ))}
          </CarouselContent>
          <CarouselPrevious className="absolute left-2 top-1/2 -translate-y-1/2 p-2 shadow-md hover:bg-transparent">
            <ChevronLeftIcon className="h-4 w-4" />
          </CarouselPrevious>
          <CarouselNext className="absolute right-2 top-1/2 -translate-y-1/2 p-2 shadow-md hover:bg-transparent">
            <ChevronRightIcon className="h-4 w-4" />
          </CarouselNext>
        </Carousel>
      </div>
      <div className="p-4">
        <div className="flex justify-between items-start">
          <div>
            <h5 className="text-lg font-semibold">{coworking.name}</h5>
          </div>
          <Badge className="bg-green-500 text-white px-2 py-1 rounded-full text-xs">Active</Badge>
        </div>
        <div className="mt-2">
          <p className="text-sm font-medium">Workspaces: {coworking.workspaces[currentWorkspaceIndex].name}</p>
        </div>
      </div>
    </Card>
  );
}

function ChevronLeftIcon(props: JSX.IntrinsicAttributes & SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="m15 18-6-6 6-6" />
    </svg>
  );
}

function ChevronRightIcon(props: JSX.IntrinsicAttributes & SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="m9 18 6-6-6-6" />
    </svg>
  );
}