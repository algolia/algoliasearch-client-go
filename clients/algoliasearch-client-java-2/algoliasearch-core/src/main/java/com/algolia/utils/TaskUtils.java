package com.algolia.utils;

import com.algolia.exceptions.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.function.IntUnaryOperator;
import java.util.function.Predicate;
import java.util.function.Supplier;

public class TaskUtils {

  public static final int DEFAULT_MAX_TRIAL = 50;
  public static final IntUnaryOperator DEFAULT_TIMEOUT = (int retries) -> {
    return Math.min(retries * 200, 5000);
  };

  public static <TResponse> void retryUntil(
    Supplier<CompletableFuture<TResponse>> func,
    Predicate<TResponse> validate,
    int maxTrial,
    IntUnaryOperator timeout
  ) throws AlgoliaRuntimeException {
    int retryCount = 0;
    while (retryCount < maxTrial) {
      try {
        TResponse resp = func.get().get();
        if (validate.test(resp)) {
          return;
        }
      } catch (InterruptedException | ExecutionException e) {
        // if the task is interrupted, just return
        return;
      }
      try {
        Thread.sleep(timeout.applyAsInt(retryCount));
      } catch (InterruptedException ignored) {
        // Restore interrupted state...
        Thread.currentThread().interrupt();
      }

      retryCount++;
    }
    throw new AlgoliaRetriesExceededException("The maximum number of trials exceeded. (" + (retryCount + 1) + "/" + maxTrial + ")");
  }
}
