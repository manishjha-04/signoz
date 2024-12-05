import getFeaturesFlags from 'api/features/getFeatureFlags';
import { REACT_QUERY_KEY } from 'constants/reactQueryKeys';
import { useQuery, UseQueryResult } from 'react-query';
import { FeatureFlagProps } from 'types/api/features/getFeaturesFlags';

const useGetFeatureFlag = (
	onSuccessHandler: (routes: FeatureFlagProps[]) => void,
	userEmail: string,
): UseQueryResult<FeatureFlagProps[], unknown> =>
	useQuery<FeatureFlagProps[]>({
		queryFn: getFeaturesFlags,
		queryKey: [REACT_QUERY_KEY.GET_FEATURES_FLAGS, userEmail],
		onSuccess: onSuccessHandler,
		retryOnMount: false,
		enabled: !!userEmail,
	});

export default useGetFeatureFlag;
