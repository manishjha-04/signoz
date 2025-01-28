/* eslint-disable react-hooks/exhaustive-deps */
import { DefaultOptionType } from 'antd/es/select';
import { getAttributesValues } from 'api/queryBuilder/getAttributesValues';
import { useQuery } from 'react-query';
import { DataTypes } from 'types/api/queryBuilder/queryAutocompleteResponse';
import { DataSource } from 'types/common/queryBuilder';

export interface Filters {
	searchText: string;
	attributeKey: string | string[];
}

export interface GetAllFiltersResponse {
	options: DefaultOptionType[];
	isFetching: boolean;
}

export function useGetAllFilters(props: Filters): GetAllFiltersResponse {
	const { searchText, attributeKey } = props;

	const { data, isLoading } = useQuery(
		['attributesValues', attributeKey, searchText],
		async () => {
			const keys = Array.isArray(attributeKey) ? attributeKey : [attributeKey];

			const responses = await Promise.all(
				keys.map((key) =>
					getAttributesValues({
						aggregateOperator: 'noop',
						dataSource: DataSource.TRACES,
						aggregateAttribute: '',
						attributeKey: key,
						searchText: searchText ?? '',
						filterAttributeKeyDataType: DataTypes.String,
						tagType: 'tag',
					}),
				),
			);

			const uniqueValues = [
				...new Set(
					responses.flatMap(
						({ payload }) => Object.values(payload || {}).find((el) => !!el) || [],
					),
				),
			];

			return uniqueValues.map((val: string) => ({
				label: val,
				value: val,
			}));
		},
	);

	return { options: data ?? [], isFetching: isLoading };
}
