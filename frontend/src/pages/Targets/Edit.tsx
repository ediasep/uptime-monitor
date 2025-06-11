import { useState, useEffect } from "react";
import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import PageMeta from "../../components/common/PageMeta";
import Input from "../../components/form/input/InputField";
import Label from "../../components/form/Label";
import Button from "../../components/ui/button/Button";
import { useParams, useNavigate } from "react-router";
import Alert from "../../components/ui/alert/Alert";
import { Modal } from "../../components/ui/modal";

export default function Edit() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [name, setName] = useState("");
  const [url, setUrl] = useState("");
  const [interval, setInterval] = useState(60);
  const [loading, setLoading] = useState(true);
  const [successMessage, setSuccessMessage] = useState("");
  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [nameError, setNameError] = useState(false);
  const [urlError, setUrlError] = useState(false);
  const [intervalError, setIntervalError] = useState(false);

  // Reset error state
  const resetErrorStates = () => {
    setNameError(false);
    setUrlError(false);
    setIntervalError(false);
  };

  // Fetch target data on mount
  useEffect(() => {
    if (!id) return;
    setLoading(true);
    fetch(`http://localhost:8080/targets/${id}`)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch target");
        return res.json();
      })
      .then((data) => {
        setName(data.name);
        setUrl(data.url);
        setInterval(data.interval);
      })
      .catch((err) => {
        alert(err.message);
      })
      .finally(() => setLoading(false));
  }, [id]);

  const handleSubmit = async () => {
    if (!id) return;

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

    const response = await fetch(`http://localhost:8080/targets/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, url, interval }),
    });

    if (response.ok) {
      resetErrorStates();
      setSuccessMessage("Target updated successfully!");
    } else {
      alert(response.statusText);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <PageMeta
        title="Uptime Monitor - Edit Target"
        description="Edit target for monitoring in Uptime Monitor."
      />

      {showConfirmModal && (
        <Modal
          isOpen={showConfirmModal}
          onClose={() => setShowConfirmModal(false)}
          className="max-w-md"
          showCloseButton={true}
          isFullscreen={false}
        >
          <div className="rounded-2xl border border-gray-200 bg-white px-5 py-16 dark:border-gray-800 dark:bg-white/[0.03] xl:px-10 xl:py-12">
            <div className="space-y-12">
              <p>Are you sure you want to delete this target?</p>
              <div className="flex gap-3">
                <Button
                  size="sm"
                  variant="outlineerror"
                  className="text-error-500"
                  onClick={() => {
                    fetch(`http://localhost:8080/targets/${id}`, {
                      method: "DELETE",
                    })
                      .then((res) => {
                        if (!res.ok) throw new Error("Failed to delete target");
                        setSuccessMessage("Target deleted successfully!");
                        setShowConfirmModal(false);
                        setTimeout(() => {
                          navigate("/"); // Navigate to target list after delete
                        }, 1000); // Optional: delay for user to see success message
                      })
                      .catch((err) => alert(err.message));
                  }}
                >
                  Delete
                </Button>
                <Button
                  size="sm"
                  onClick={() => setShowConfirmModal(false)}
                >
                  Cancel
                </Button>
                </div>
            </div>
          </div>
        </Modal>
      )}

      <PageBreadcrumb pageTitle="Edit Target" />
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
            <Label htmlFor="input-name">Name</Label>
            <Input
              type="text"
              id="input-name"
              placeholder="Target Name"
              value={name}
              error={nameError}
              hint={nameError ? "Target Name is Required." : ""}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div>
            <Label htmlFor="input-url">URL</Label>
            <Input
              type="text"
              id="input-url"
              placeholder="Target URL"
              value={url}
              error={urlError}
              hint={urlError ? "Target URL is Required." : ""}
              onChange={(e) => setUrl(e.target.value)}
            />
          </div>
          <div>
            <Label htmlFor="input-interval">Interval</Label>
            <Input
              type="number"
              id="input-interval"
              placeholder="Ping Interval"
              value={interval}
              error={intervalError}
              hint={intervalError ? "Interval is Required and must be a positive number." : ""}
              onChange={(e) => setInterval(parseInt(e.target.value, 10))}
            />
          </div>

          <div className="gap-3 pb-3 flex flex-wrap">
            <Button size="sm" onClick={handleSubmit}>
              Update Target
            </Button>
            
            <Button size="sm" variant="texterror" className="text-error-500" onClick={() => setShowConfirmModal(true)}>
              Delete Target
            </Button>
          </div>

        </div>
      </div>
    </div>
  );
}
