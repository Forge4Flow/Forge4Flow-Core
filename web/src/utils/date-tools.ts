// ** Date-FNS Imports
import { formatDistanceToNow, format } from 'date-fns'

export function convertDate(dateString: string): string {
  const date = new Date(dateString)
  const formattedDate = format(date, 'MMM do yyyy')
  const timeAgo = formatDistanceToNow(date, { addSuffix: true })

  return `${formattedDate} ${timeAgo}`
}
