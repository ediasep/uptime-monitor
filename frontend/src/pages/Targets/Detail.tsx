import { useParams } from "react-router";
import PageBreadcrumb from "../../components/common/PageBreadCrumb";
import PageMeta from "../../components/common/PageMeta";
import { useEffect, useState } from "react";
import ComponentCard from "../../components/common/ComponentCard";
import BarChartOne from "../../components/charts/bar/BarChartOne";

export default function Detail() {
    const { id } = useParams();
    const [name, setName] = useState("");
    const [url, setUrl] = useState("");
    const [loading, setLoading] = useState(true);
    const [categories, setCategories] = useState<string[]>([]);
    const [series, setSeries] = useState<{ name: string; data: number[] }[]>([]);

    useEffect(() => {
        if (!id) return;

        const fetchData = async () => {
            setLoading(true);
            try {
                const [uptimeRes, targetRes] = await Promise.all([
                    fetch(`http://localhost:8080/targets/${id}/uptime/daily`),
                    fetch(`http://localhost:8080/targets/${id}`)
                ]);

                if (!uptimeRes.ok) throw new Error("Failed to fetch uptime data");
                if (!targetRes.ok) throw new Error("Failed to fetch target");

                const uptimeData = await uptimeRes.json();
                const targetData = await targetRes.json();

                setCategories(
                    uptimeData.map((item: any) =>
                        new Date(item.date).toLocaleDateString(undefined, {
                            year: "numeric",
                            month: "short",
                            day: "numeric",
                        })
                    )
                );
                setSeries([
                    {
                        name: "Uptime (%)",
                        data: uptimeData.map((item: any) =>
                          Math.round(item.uptime_percentage * 100) / 100
                        ),
                    },
                ]);
                setName(targetData.name);
                setUrl(targetData.url);
            } catch (err: any) {
                alert(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [id]);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <PageMeta
                title="Uptime Monitoring - Target Detail"
                description="View detailed information about a specific target."
            />
            <PageBreadcrumb pageTitle={name} pageDescription={url} />
            
            <ComponentCard title="Last 7 Days Uptime">
                <BarChartOne categories={categories} series={series} />
            </ComponentCard>
        </div>
    )
}