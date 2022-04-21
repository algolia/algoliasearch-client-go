package com.algolia;

import com.algolia.exceptions.AlgoliaRuntimeException;
import java.util.List;
import java.util.Map;

/**
 * Callback for asynchronous API call.
 *
 * @param <T> The return type
 */
public interface ApiCallback<T> {
  /**
   * This is called when the API call fails.
   *
   * @param e The exception causing the failure
   * @param statusCode Status code of the response if available, otherwise it would be 0
   * @param responseHeaders Headers of the response if available, otherwise it would be null
   */
  void onFailure(
    AlgoliaRuntimeException e,
    int statusCode,
    Map<String, List<String>> responseHeaders
  );

  /**
   * This is called when the API call succeeded.
   *
   * @param result The result deserialized from response
   * @param statusCode Status code of the response
   * @param responseHeaders Headers of the response
   */
  void onSuccess(
    T result,
    int statusCode,
    Map<String, List<String>> responseHeaders
  );

  /**
   * This is called when the API upload processing.
   *
   * @param bytesWritten bytes Written
   * @param contentLength content length of request body
   * @param done write end
   */
  void onUploadProgress(long bytesWritten, long contentLength, boolean done);

  /**
   * This is called when the API download processing.
   *
   * @param bytesRead bytes Read
   * @param contentLength content length of the response
   * @param done Read end
   */
  void onDownloadProgress(long bytesRead, long contentLength, boolean done);
}
