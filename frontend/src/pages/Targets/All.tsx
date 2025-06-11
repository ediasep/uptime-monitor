import { useNavigate } from "react-router";
import PageMeta from "../../components/common/PageMeta";
import TargetTable from "../../components/tables/TargetTable";
import Button from "../../components/ui/button/Button";

import {
  PlusIcon
} from "../../icons";
import PageBreadcrumb from "../../components/common/PageBreadCrumb";

export default function All() {

  const navigate = useNavigate();

  const handleClick = () => {
    navigate("/targets/create");
  };

  return (
    <div>
      <PageMeta
        title="Uptime Monitoring - All Targets"
        description="View and manage all your uptime monitoring targets."
      />

      <PageBreadcrumb pageTitle="All Target" />
      <div className="rounded-2xl border border-gray-200 bg-white px-5 py-7 dark:border-gray-800 dark:bg-white/[0.03] xl:px-10 xl:py-12">
        
        <Button 
            size="sm" 
            variant="outline" 
            className="mb-6" 
            startIcon={<PlusIcon />} 
            onClick={handleClick}
        >
           Add Target
        </Button>

        <TargetTable />
      </div>
    </div>
  );
}
