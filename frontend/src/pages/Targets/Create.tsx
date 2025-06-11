import { useState } from "react";
import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import PageMeta from "../../components/common/PageMeta";
import Input from "../../components/form/input/InputField";
import Label from "../../components/form/Label";
import Button from "../../components/ui/button/Button";
import Alert from "../../components/ui/alert/Alert";

export default function Create() {
  const [name, setName] = useState("");
  const [url, setUrl] = useState("");
  const [interval, setInterval] = useState(60);
  const [successMessage, setSuccessMessage] = useState("");
  const [nameError, setNameError] = useState(false);
  const [urlError, setUrlError] = useState(false);
  const [intervalError, setIntervalError] = useState(false);

  // Reset error state
  const resetErrorStates = () => {
    setNameError(false);
    setUrlError(false);
    setIntervalError(false);
  };


  const handleSubmit = async () => {
    // Reset error states
    resetErrorStates();

    if( !name || !url || !interval || isNaN(interval) || interval <= 0) {
      if (!name) {
        setNameError(true);
      }
      if (!url) {
        setUrlError(true);
      }
      if (!interval || isNaN(interval) || interval <= 0) {
        setIntervalError(true);
      }

      return;
    }

    const response = await fetch("http://localhost:8080/targets", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, url, interval }),
    });

    if (response.ok) {
      resetErrorStates();
      setSuccessMessage("Target created successfully!");
    } else {
      alert(response.statusText);
    }
  };

  return (
    <div>
      <PageMeta
        title=" Uptime Monitor - Create Target"
        description="Create a new target for monitoring in Uptime Monitor."
      />
      <PageBreadcrumb pageTitle="Create Target" />
      <div className="rounded-2xl border border-gray-200 bg-white px-5 py-7 dark:border-gray-800 dark:bg-white/[0.03] xl:px-10 xl:py-12">
        <div className="space-y-6">
          {successMessage && (
            <Alert
              variant="success"
              title="Success"
              message={successMessage}
              showLink={true}
              linkHref="/"
              linkText="Show Target List"
            />
          )}

          <div>
            <Label htmlFor="input">Target Name</Label>
            <Input type="text" id="input" placeholder="Target Name" value={name} error={nameError} hint={nameError ? "Target Name is Required." : ""} onChange={(e) => setName(e.target.value)}/>
          </div>
          <div>
            <Label htmlFor="input">URL</Label>
            <Input type="text" id="input" placeholder="Target URL" value={url} error={urlError} hint={urlError ? "Target URL is Required." : ""} onChange={(e) => setUrl(e.target.value)}/>
          </div>
          <div>
            <Label htmlFor="input">Interval</Label>
            <Input type="text" id="input" placeholder="Ping Interval" value={interval} error={intervalError} hint={intervalError ? "Inverval is Required and must be a positive number." : ""} onChange={(e) => setInterval(parseInt(e.target.value, 10))}/>
          </div>

          <Button size="sm" className="place-items-end" onClick={handleSubmit}>
            Create Target
          </Button>
        </div>

      </div>
    </div>
  );
}
