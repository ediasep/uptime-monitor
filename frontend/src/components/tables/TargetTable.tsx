import {
  Table,
  TableBody,
  TableCell,
  TableHeader,
  TableRow,
} from "../ui/table";


import { useEffect, useState } from "react";
import { PencilIcon } from "../../icons";
import { useNavigate } from "react-router";
import Badge from "../ui/badge/Badge";

interface Target {
  id: string;
  name: string;
  url: string;
  interval: number;
  last_status: string; // "UP" or "DOWN"
  last_checked_at: string; // ISO date string
  created_at: string; // ISO date string
}

export default function TargetTable() {
  const [targets, setTargets] = useState<Target[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost:8080/targets")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch targets");
        }
        return response.json();
      })
      .then((data) => setTargets(data))
      .catch((error) => console.error("Error:", error));
  }, []); // run once on mount

  const onEdit = (id: string) => {
    navigate("/targets/edit/" + id);
  };

  return (
    <div className="overflow-hidden rounded-xl border border-gray-200 bg-white dark:border-white/[0.05] dark:bg-white/[0.03]">
      <div className="max-w-full overflow-x-auto">
        <Table>
          {/* Table Header */}
          <TableHeader className="border-b border-gray-100 dark:border-white/[0.05]">
            <TableRow>
              <TableCell
                isHeader
                className="px-5 py-3 font-medium text-gray-500 text-start text-theme-xs dark:text-gray-400"
              >
                Name
              </TableCell>
              <TableCell
                isHeader
                className="px-5 py-3 font-medium text-gray-500 text-start text-theme-xs dark:text-gray-400"
              >
                URL
              </TableCell>
              <TableCell
                isHeader
                className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400"
              >
                Last Status
              </TableCell>
              <TableCell
                isHeader
                className="px-5 py-3 font-medium text-gray-500 text-start text-theme-xs dark:text-gray-400"
              >
                Last Checked At
              </TableCell>
              <TableCell
                isHeader
                className="px-5 py-3 font-medium text-gray-500 text-center text-theme-xs dark:text-gray-400"
              >
                Action
              </TableCell>
            </TableRow>
          </TableHeader>

          {/* Table Body */}
          <TableBody className="divide-y divide-gray-100 dark:divide-white/[0.05]">
            {targets.map((target) => (
              <TableRow key={target.id}>
                <TableCell className="px-4 py-3 text-gray-500 text-start text-theme-sm dark:text-gray-400">
                  {target.name}
                </TableCell>
                <TableCell className="px-4 py-3 text-gray-500 text-start text-theme-sm dark:text-gray-400">
                  {target.url}
                </TableCell>
                <TableCell className="px-4 py-3 text-gray-500 text-theme-sm dark:text-gray-400 text-center">
                  {target.last_status === "UP" ? (
                    <Badge children="UP" color="success" variant="solid" />
                  ) : (
                    <Badge children="DOWN" color="error" variant="solid" />
                  )}
                </TableCell>
                <TableCell className="px-4 py-3 text-gray-500 text-theme-sm dark:text-gray-400">
                  <span className="text-gray-400">
                    {target.last_checked_at
                      ? new Date(target.last_checked_at).toLocaleString(
                          undefined,
                          {
                            year: "numeric",
                            month: "short",
                            day: "numeric",
                            hour: "2-digit",
                            minute: "2-digit",
                            second: "2-digit",
                          }
                        )
                      : "Never checked"}
                  </span>
                </TableCell>
                <TableCell className="px-4 py-3 text-gray-500 text-theme-sm dark:text-gray-400 text-center align-middle">
                  <span
                    className="flex justify-center items-center h-full cursor-pointer"
                    onClick={() => onEdit(target.id)}
                  >
                    <PencilIcon width={21} height={21} />
                  </span>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
