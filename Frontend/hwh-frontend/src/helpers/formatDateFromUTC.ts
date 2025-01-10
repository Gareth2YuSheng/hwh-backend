import { formatDistance } from 'date-fns';

const formatDateFromUTC = (date : string) => {
  const initialDate = Date.parse(date);
  const today = new Date();
  const options = { addSuffix: true };
  return formatDistance(initialDate, today, options);
}

export default formatDateFromUTC;