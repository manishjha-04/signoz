export interface UpdateProfileProps {
	reasons_for_interest_in_signoz: string;
	uses_otel: boolean;
	has_existing_observability_tool: boolean;
	existing_observability_tool: string;
	logs_scale_per_day_in_gb: number;
	number_of_services: number;
	number_of_hosts: number;
	where_did_you_discover_signoz: string;
}
